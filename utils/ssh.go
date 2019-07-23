package utils

//https://play.golang.org/p/VgFKqjFlUGu

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

var rawKey = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACC0CmZM+kb44TryzZRA9T8tC8DacC8KfB8xTCiFhNaR8wAAAJhNahWvTWoV
rwAAAAtzc2gtZWQyNTUxOQAAACC0CmZM+kb44TryzZRA9T8tC8DacC8KfB8xTCiFhNaR8w
AAAEBrGM/hXTiyPUG2hDoqnz2g123+15rcbVtBU1J1o5YBwLQKZkz6RvjhOvLNlED1Py0L
wNpwLwp8HzFMKIWE1pHzAAAAD21hbnVAbnl4Mi5sb2NhbAECAwQFBg==
-----END OPENSSH PRIVATE KEY-----`


func main() {
	priv, err := ssh.ParseRawPrivateKey([]byte(rawKey))
	fmt.Println(err)
	signer, _ := ssh.NewSignerFromKey(priv)
	pubKeyBytes := ssh.MarshalAuthorizedKey(signer.PublicKey())
	fmt.Println(string(pubKeyBytes))
}

/*
def get_private_keys():
    """Find SSH keys in standard folder."""
    key_formats = [RSAKey, ECDSAKey, Ed25519Key]

    ssh_folder = os.path.expanduser('~/.ssh')

    available_private_keys = []
    if os.path.isdir(ssh_folder):
        for key in os.listdir(ssh_folder):
            key_file = os.path.join(ssh_folder, key)
            if not os.path.isfile(key_file):
                continue
            for key_format in key_formats:
                try:
                    parsed_key = key_format.from_private_key_file(key_file)
                    key_details = {
                        'filename': key,
                        'format': parsed_key.get_name(),
                        'bits': parsed_key.get_bits(),
                        'fingerprint': parsed_key.get_fingerprint().hex()
                    }
                    available_private_keys.append(key_details)
                except (SSHException, UnicodeDecodeError, IsADirectoryError):
                    continue
                except OSError as e:
                    if e.errno == errno.ENXIO:
                        # when key_file is a (ControlPath) socket
                        continue
                    else:
                        raise

    return available_private_keys

 */
