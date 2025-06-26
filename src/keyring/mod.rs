use bip39::Mnemonic;

pub struct Keyring {
    pub storage_path: String,
    pub mnemonic: String,
}

impl Keyring {
    pub fn new(storage_path: String) -> Self {
        Self {
            storage_path,
            mnemonic: String::new(),
        }
    }

    pub fn create_new_wallet(&mut self, word_count: usize) -> Result<String, String> {
        if word_count != 12 && word_count != 15 && word_count != 18 && word_count != 21 && word_count != 24 {
            return Err("Invalid word count. Must be one of: 12, 15, 18, 21, or 24.".to_string());
        }
        
        match Self::generate(word_count) {
            Ok(mnemonic) => {
                self.mnemonic = mnemonic;
                Ok(self.mnemonic.clone())
            },
            Err(e) => Err(format!("Failed to generate mnemonic: {}", e)),
        }
    }

    fn generate(word_count: usize) -> Result<String, String> {
        let mnemonic = Mnemonic::generate(word_count).unwrap();
        Ok(mnemonic.to_string())
    }
}
