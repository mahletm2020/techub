import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client/core';
import { defineNuxtPlugin } from '#app';

export default defineNuxtPlugin((nuxtApp) => {
  const httpLink = new HttpLink({
    uri: 'http://localhost:9702/v1/graphql',
    headers: {
'x-hasura-admin-secret': process.env.HASURA_ADMIN_SECRET || '',    },
  });

  const apolloClient = new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache(),
  });

  // Provide Apollo Client globally
  nuxtApp.provide('apollo', apolloClient);
});
