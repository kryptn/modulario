import React, { Component } from 'react'
import Fetch from 'react-fetch-component'
import {
  BrowserRouter as Router,
  Route,
  Link
} from 'react-router-dom'

const Home = () => (
  <div>
    <h2>Home</h2>
  </div>
)

const About = () => (
  <div>
    <h2>About</h2>
  </div>
)

const Topic = ({ match }) => (
  <div>
    <h3>{match.params.topicId}</h3>
  </div>
)

const Topics = ({ match }) => (
  <div>
    <h2>Topics</h2>
    <ul>
      <li>
        <Link to={`${match.url}/rendering`}>
          Rendering with React
        </Link>
      </li>
      <li>
        <Link to={`${match.url}/components`}>
          Components
        </Link>
      </li>
      <li>
        <Link to={`${match.url}/props-v-state`}>
          Props v. State
        </Link>
      </li>
    </ul>

    <Route path={`${match.url}/:topicId`} component={Topic}/>
    <Route exact path={match.url} render={() => (
      <h3>Please select a topic.</h3>
    )}/>
  </div>
)


const CreatePost = ({ match }) => (
  <div> words </div>
)

const ViewPost = ({ match }) => (
  <div>
    <Fetch url={`http://localhost:5000/api/v1/view/${match.params.postId}`}>
      {({ loading, data, error }) => (
        <div>
          {loading && <span>Loading...</span>}
          {data && <div><pre>{JSON.stringify(data, null, 2) }</pre></div>}
        </div>
      )}
      </Fetch>
    </div>
)


class App extends Component {
  render() {
    return (
      <Router>
         <div>
           <ul>
             <li><Link to="/">Home</Link></li>
             <li><Link to="/view">View Post</Link></li>
             <li><Link to="/create">Create Post</Link></li>
           </ul>

           <hr/>

           <Route exact path="/" component={Home}/>
           <Route path="/create" component={CreatePost}/>
           <Route path="/view/:postId" component={ViewPost}/>

         </div>
       </Router>
    );
  }
}

export default App;
