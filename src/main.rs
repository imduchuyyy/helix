
mod keyring;

fn main() {
    let keyring_service = keyring::Keyring::new("/.crypto-lite".to_string());
}