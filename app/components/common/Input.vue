<template>
  <div class="space-y-2">
    <label v-if="label" :for="id" class="block text-sm font-medium text-foreground">
      {{ label }}
      <span v-if="required" class="text-destructive">*</span>
    </label>
    <input
      :id="id"
      :type="type"
      :value="modelValue"
      :placeholder="placeholder"
      :required="required"
      :disabled="disabled"
      :class="inputClasses"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
    />
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

defineEmits<{
  'update:modelValue': [value: string]
}>()

const inputClasses = computed(() => {
  const base = 'input'
  const errorClass = props.error ? 'border-destructive focus-visible:ring-destructive' : ''
  return `${base} ${errorClass}`
})
</script>
