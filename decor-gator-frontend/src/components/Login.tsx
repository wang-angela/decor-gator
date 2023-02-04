import { useNavigate } from 'react-router-dom';
import LoginField from './LoginField'

function Login() {
    return (
    <div className="App">
      <span className="heading">Decor Gator</span>
      <LoginField />
    </div>
    );
}

export default Login;