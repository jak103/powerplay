import { defineStore } from 'pinia';

export const useChannelStore = defineStore('channel', {
state: () => ({
    channels: [
        {
        channel_name: 'Trash Pandas',
        channel_id: 'example1',
        channel_image: '/_nuxt/assets/trash-panda.webp',
        missed_chats: 1,
        type: 'team'
      },
      {
        channel_name: 'District 5',
        channel_id: 'example2',
        channel_image: '/_nuxt/assets/District-5.webp',
        missed_chats: 16,
        type: 'team'
      },
      {
        channel_name: 'C-League',
        channel_id: 'example3',
        channel_image: null,
        missed_chats: 0,
        type: 'league'
      },
      {
        channel_name: 'D-League',
        channel_id: 'example4',
        channel_image: null,
        missed_chats: 0,
        type: 'league'
      },
      {
        channel_name: 'C-League Subs',
        channel_id: 'example5',
        channel_image: null,
        missed_chats: 3,
        type: 'channel'
      },
      {
        channel_name: 'D-League Subs',
        channel_id: 'example6',
        channel_image: null,
        missed_chats: 0,
        type: 'channel'
      },
      {
        channel_name: 'News & Announcements',
        channel_id: 'example7',
        channel_image: null,
        missed_chats: 1,
        type: 'channel'
      },
      {
        channel_name: 'Ricky Bobby',
        channel_id: 'example8',
        channel_image: '/_nuxt/assets/ProfilePic.webp',
        missed_chats: 0,
        type: 'dm'
      },
      {
        channel_name: 'D-League Subs',
        channel_id: 'example9',
        channel_image: null,
        missed_chats: 3,
        type: 'dm'
      },
      {
        channel_name: 'Steve?? ',
        channel_id: 'example10',
        channel_image: null,
        missed_chats: 0,
        type: 'dm'
      },
      {
        channel_name: 'Fabio Lanzoni',
        channel_id: 'example11',
        channel_image: null,
        missed_chats: 0,
        type: 'dm'
      }
    ]
  }),
  actions: {
    addChannel(channel: {
      channel_name: string;
      channel_id: string;
      channel_image: string | null;
      missed_chats: number;
      type: string;
    }) {
      this.channels.push(channel);
    },
    removeChannel(channelId: string) {
      this.channels = this.channels.filter(channel => channel.channel_id !== channelId);
    },
    getChannelById(channelId: string) {
      return this.channels.find(channel => channel.channel_id === channelId);
    }
  }
});
