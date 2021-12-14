import {Buffer} from 'buffer/';
const  crypto =  require("crypto-browserify");

export const getBase64 = async (file: any) => {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => resolve(reader.result);
        reader.onerror = error => reject(error);
  });
}


export const encrypt = (plaintext: string, key: string): string => {

    let buffer = new Buffer(plaintext);
    let encrypted = crypto.privateEncrypt(key, buffer);

    return encrypted.toString('base64');
 }

export const decrypt = (cypher: string, key: string): string => {

    let buffer = Buffer.from(cypher, 'base64');
    let plaintext = crypto.publicDecrypt(key, buffer);

    return plaintext.toString('utf8')
}
