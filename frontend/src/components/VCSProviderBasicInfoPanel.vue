<template>
  <div class="textlabel">
    {{ $t("version-control.setting.add-git-provider.choose") }}
    <span class="text-red-600">*</span>
  </div>
  <div class="pt-4 radio-set-row">
    <div class="radio space-x-2">
      <input
        v-model="config.type"
        name="Self-host GitLab"
        tabindex="-1"
        type="radio"
        class="btn"
        value="GITLAB_SELF_HOST"
      />
      <img class="h-6 w-auto" src="../assets/gitlab-logo.svg" />
      <label class="whitespace-nowrap"
        >{{
          $t("version-control.setting.add-git-provider.gitlab-self-host-ce-ee")
        }}
      </label>
    </div>
  </div>
  <div class="mt-4 relative">
    <div class="relative flex justify-start">
      <span class="pr-2 bg-white text-xs text-control-light">
        {{ $t("common.coming-later") }}
      </span>
    </div>
  </div>
  <div class="mt-2 flex flex-row itmes-center space-x-4 text-xs">
    <div class="flex flex-row space-x-2 items-center text-control">
      <div class="h-5 w-5">
        <img src="../assets/gitlab-logo.svg" />
      </div>
      <label class="whitespace-nowrap">GitLab.com </label>
    </div>
    <div class="flex flex-row space-x-2 items-center text-control">
      <div class="h-5 w-5">
        <img src="../assets/github-logo.svg" />
      </div>
      <label class="whitespace-nowrap">GitHub Enterprise </label>
    </div>
    <div class="flex flex-row space-x-2 items-center text-control">
      <div class="h-5 w-5">
        <img src="../assets/github-logo.svg" />
      </div>
      <label class="whitespace-nowrap">GitHub.com </label>
    </div>
  </div>
  <div class="mt-6 pt-6 border-t border-block-border textlabel">
    {{ instanceUrlLabel }} <span class="text-red-600">*</span>
  </div>
  <p class="mt-1 textinfolabel">
    {{
      $t(
        "version-control.setting.add-git-provider.basic-info.gitlab-instance-url-label"
      )
    }}
  </p>
  <BBTextField
    class="mt-2 w-full"
    :placeholder="'https://gitlab.example.com'"
    :value="config.instanceUrl"
    @input="changeUrl($event.target.value)"
  />
  <p v-if="state.showUrlError" class="mt-2 text-sm text-error">
    {{
      $t(
        "version-control.setting.add-git-provider.basic-info.instance-url-error"
      )
    }}
  </p>
  <div class="mt-4 textlabel">
    {{ $t("version-control.setting.add-git-provider.basic-info.display-name") }}
  </div>
  <p class="mt-1 textinfolabel">
    {{
      $t(
        "version-control.setting.add-git-provider.basic-info.display-name-label"
      )
    }}
  </p>
  <BBTextField
    class="mt-2 w-full"
    :placeholder="namePlaceholder"
    :value="config.name"
    @input="config.name = $event.target.value"
  />
</template>

<script lang="ts">
import { computed, onUnmounted, PropType, reactive } from "vue";
import isEmpty from "lodash-es/isEmpty";
import { TEXT_VALIDATION_DELAY, VCSConfig } from "../types";
import { isUrl } from "../utils";
import { useI18n } from "vue-i18n";

interface LocalState {
  urlValidationTimer?: ReturnType<typeof setTimeout>;
  showUrlError: boolean;
}

export default {
  name: "VCSProviderBasicInfoPanel",
  props: {
    config: {
      required: true,
      type: Object as PropType<VCSConfig>,
    },
  },
  setup(props) {
    const { t } = useI18n();
    const state = reactive<LocalState>({
      showUrlError:
        !isEmpty(props.config.instanceUrl) && !isUrl(props.config.instanceUrl),
    });

    onUnmounted(() => {
      if (state.urlValidationTimer) {
        clearInterval(state.urlValidationTimer);
      }
    });

    const namePlaceholder = computed((): string => {
      if (props.config.type == "GITLAB_SELF_HOST") {
        return t("version-control.setting.add-git-provider.gitlab-self-host");
      }
      return "";
    });

    const instanceUrlLabel = computed((): string => {
      if (props.config.type == "GITLAB_SELF_HOST") {
        return t(
          "version-control.setting.add-git-provider.basic-info.gitlab-instance-url"
        );
      }
      return "";
    });

    const changeUrl = (value: string) => {
      props.config.instanceUrl = value;

      if (state.urlValidationTimer) {
        clearInterval(state.urlValidationTimer);
      }
      // If text becomes valid, we immediately clear the error.
      // otherwise, we delay TEXT_VALIDATION_DELAY to do the validation in case there is continous keystroke.
      if (isUrl(props.config.instanceUrl)) {
        state.showUrlError = false;
      } else {
        state.urlValidationTimer = setTimeout(() => {
          // If error is already displayed, we hide the error only if there is valid input.
          // Otherwise, we hide the error if input is either empty or valid.
          if (state.showUrlError) {
            state.showUrlError = !isUrl(props.config.instanceUrl);
          } else {
            state.showUrlError =
              !isEmpty(props.config.instanceUrl) &&
              !isUrl(props.config.instanceUrl);
          }
        }, TEXT_VALIDATION_DELAY);
      }
    };

    return { state, namePlaceholder, instanceUrlLabel, changeUrl };
  },
};
</script>
