import type { Meta, StoryObj } from '@storybook/vue3';

import ScheduleList from './ScheduleList.vue';

const meta: Meta<typeof ScheduleList> = {
  component: ScheduleList,
};

export default meta;
type Story = StoryObj<typeof ScheduleList>;

export const Primary: Story = {
  render: (args) => ({
    components: { ScheduleList },
    setup() {
      return { args };
    },
    template: '<ScheduleList v-bind="args" />',
  }),
  args: {
    color: ''
  },
};
