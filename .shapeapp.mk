include $(PATH_TO_ROOT)/.config.mk

PATH_TO_SHAPEAPPMK=$(PATH_TO_ROOT)/vendor/shapeappmk
include $(PATH_TO_SHAPEAPPMK)/components/molecules/make/default.mk
include $(PATH_TO_SHAPEAPPMK)/components/atoms/asdf/reshim.mk
include $(PATH_TO_SHAPEAPPMK)/components/atoms/golang/tool.mk
include $(PATH_TO_SHAPEAPPMK)/components/atoms/text-template-cli/base.mk
include $(PATH_TO_SHAPEAPPMK)/default/app_stage.mk
include $(PATH_TO_SHAPEAPPMK)/default/ports.mk
include $(PATH_TO_SHAPEAPPMK)/default/directories.mk
