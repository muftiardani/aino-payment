<template>
  <div class="space-y-2">
    <label v-if="label" :for="id" class="block text-sm font-medium text-foreground">
      {{ label }}
      <span v-if="required" class="text-destructive">*</span>
    </label>
    <div class="relative">
      <div
        v-if="$slots.prefix"
        class="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground"
      >
        <slot name="prefix" />
      </div>
      <input
        :id="id"
        :type="type"
        :value="modelValue"
        :placeholder="placeholder"
        :required="required"
        :disabled="disabled"
        :class="[inputClasses, { 'pl-10': $slots.prefix, 'pr-10': $slots.suffix }]"
        @input="handleInput"
        @blur="handleBlur"
      />
      <div
        v-if="$slots.suffix"
        class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground"
      >
        <slot name="suffix" />
      </div>
    </div>
    <p v-if="error" class="text-sm text-destructive">{{ error }}</p>
    <p v-else-if="hint" class="text-sm text-muted-foreground">{{ hint }}</p>
  </div>
</template>

<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    id?: string
    label?: string
    type?: string
    modelValue?: string | number
    placeholder?: string
    required?: boolean
    disabled?: boolean
    error?: string
    hint?: string
  }>(),
  {
    type: 'text',
    required: false,
    disabled: false,
  }
)

const emit = defineEmits<{
  'update:modelValue': [value: string]
  blur: [event: FocusEvent]
}>()

const handleInput = (event: Event) => {
  const value = (event.target as HTMLInputElement).value
  emit('update:modelValue', value)
}

const handleBlur = (event: FocusEvent) => {
  emit('blur', event)
}

const inputClasses = computed(() => {
  const base = 'input'
  const errorClass = props.error ? 'border-destructive focus-visible:ring-destructive' : ''
  return `${base} ${errorClass}`
})
</script>
