include $(PATH_TO_ROOT)/.config.mk

PATH_TO_SHAPEAPPMK=$(PATH_TO_ROOT)/vendor/shapeappmk
include $(PATH_TO_SHAPEAPPMK)/make/debug.mk
include $(PATH_TO_SHAPEAPPMK)/make/shell.mk
include $(PATH_TO_SHAPEAPPMK)/asdf/reshim.mk
include $(PATH_TO_SHAPEAPPMK)/text-template-cli/base.mk
include $(PATH_TO_SHAPEAPPMK)/default/app_stage.mk
include $(PATH_TO_SHAPEAPPMK)/default/ports.mk
include $(PATH_TO_SHAPEAPPMK)/default/directories.mk
