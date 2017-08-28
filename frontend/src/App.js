import React, { Component } from 'react'
import Fetch from 'react-fetch-component'
import {
  BrowserRouter as Router,
  Route,
  Redirect,
  Link
} from 'react-router-dom'

const Home = () => (
  <div>
    <h2>Home</h2>
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
          {error && <span>Post not found</span>}
        </div>
      )}
      </Fetch>
    </div>
)

const VisitPost = ({ match }) => (
  window.location = `http://localhost:5000/api/v1/${match.params.postId}`
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
           <Route path="/:postId" component={VisitPost}/>

         </div>
       </Router>
    );
  }
}

export default App;
