import { createStackNavigator, createSwitchNavigator } from 'react-navigation'
import { Platform } from 'react-native'
import List from '../Screens/Chats/List'
import Detail from '../Screens/Chats/Detail'
import Add from '../Screens/Chats/Add'
import ChatSettingsNavigator from './ChatSettingsNavigator'

const handleNavigationsOptions = ({ navigation, screenProps }) => {
  let tabBarVisible = false
  const allowedViews = ['chats/list', 'chats/home']

  if ((Platform.OS === 'web' && allowedViews.indexOf(navigation.state.routeName) !== -1) ||
    (Platform.OS !== 'web' &&
      allowedViews.indexOf(navigation.state.routes[navigation.state.index].routeName) !== -1)) {
    tabBarVisible = true
  }

  return {
    tabBarVisible,
  }
}

export const SplitSideChatNavigator = createStackNavigator(
  {
    'chats/list': List,
  },
  {
    initialRouteName: 'chats/list',
    navigationOptions: handleNavigationsOptions,
  }
)

export const SubviewsChatDetailsNavigator = createStackNavigator(
  {
    'chats/detail': Detail,
    'chats/settings': {
      screen: ChatSettingsNavigator,
      navigationOptions: () => ({
        header: null,
      }),
    },
  },
  {
    initialRouteName: 'chats/detail',
    navigationOptions: handleNavigationsOptions,
  }
)

export const SubviewsChatAddNavigator = createStackNavigator(
  {
    'chats/add': Add,
  },
  {
    initialRouteName: 'chats/add',
    navigationOptions: handleNavigationsOptions,
  }
)

export const SubviewsChatNavigator = createSwitchNavigator({
  SubviewsChatDetailsNavigator,
  SubviewsChatAddNavigator,
})

export default createStackNavigator(
  {
    'chats/home': SplitSideChatNavigator,
    'chats/subviews': SubviewsChatNavigator,
  },
  {
    initialRouteName: 'chats/home',
    navigationOptions: handleNavigationsOptions,
    headerMode: 'none',
  }
)
