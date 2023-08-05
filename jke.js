const { log } = console;
/**
 * L fdph, L vdz, L frqtxhuhg，每个字母向前移动 3个位置，能得到什么字符串？
 */

const str = "L fdph, L vdz, L frqtxhuhg";

const strAdd = (...strList) => {
  return String.fromCodePoint(
    strList.reduce(
      (acc, curr) =>
        (acc += typeof curr === "number" ? curr : curr.codePointAt()),
      0
    )
  );
};

const vigenereCiper = (str, keyword) => {
  const klen = keyword.length;
  const acode = "a".codePointAt();
  const zcode = "z".codePointAt();
  const Acode = "A".codePointAt();
  const Zcode = "Z".codePointAt();

  return str.split("").reduce((acc, curr, idx) => {
    // 对当前的字符进行加密
    const codePoint = curr.codePointAt();
    const currentKeystr = keyword[idx % klen];

    if (codePoint >= acode && codePoint <= zcode) {
      const str = strAdd(curr, currentKeystr.codePointAt() - acode);
      acc += strAdd((str.codePointAt() - acode) % 26, acode);
    } else if (codePoint >= Acode && codePoint <= Zcode) {
      const str = strAdd(curr, currentKeystr.codePointAt() - Acode);
      acc += strAdd((str.codePointAt() - Acode) % 26, Acode);
    } else {
      acc += curr;
    }

    return acc;
  }, "");
};

log(vigenereCiper("ISOITEUIWUIZNSROCNKFD", "GOLANG"));
log(vigenereCiper("ISOITEUIWUIZNSROCNKFD", "GOLANG"));
log(vigenereCiper("abcdefghijklmnopqrstuvwxyz", "ddd"));

/**
  注意：维吉尼亚加密算法对英文字母进行加密，非英文字母的字符将会保持不变。
  密钥key可以是任意的字符串，它会重复使用直至与明文长度相同。在实际应用中，
  密钥的选择是非常重要的，较长的密钥通常会增加加密的安全性。
 */
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
const vigenereEncrypt = (plaintext = "", key = "") => {
  plaintext = plaintext.toUpperCase();
  key = key.toUpperCase();

  let ciphertext = "";

  for (let i = 0; i < plaintext.length; i++) {
    const plainChar = plaintext[i];
    const keyChar = key[i % key.length];

    if (alphabet.includes(plainChar)) {
      const plainIndex = alphabet.indexOf(plainChar);
      const keyIndex = alphabet.indexOf(keyChar);

      const entrytedIndex = (plainIndex + keyIndex) % 26;
      ciphertext += alphabet[entrytedIndex];
    } else {
      ciphertext += plainChar;
    }
  }
  return ciphertext;
};

const vigenereDecrypt = (ciphertext, key) => {
  ciphertext = ciphertext.toUpperCase();
  key = key.toUpperCase();

  let plaintext = "";

  for (let i = 0; i < ciphertext.length; i++) {
    const cipherChar = ciphertext[i];
    const keyChar = key[i % key.length];

    if (alphabet.includes(cipherChar)) {
      const cipherIndex = alphabet.indexOf(cipherChar);
      const keyIndex = alphabet.indexOf(keyChar);

      const decryptedIndex = Math.abs(cipherIndex - keyIndex) % 26;
      plaintext += alphabet[decryptedIndex];
    } else {
      plaintext += cipherChar;
    }
  }

  return plaintext;
};

log(vigenereEncrypt("ISOITEUIWUIZNSROCNKFD", "GOLANG"));
log(vigenereDecrypt("IGZIGKAWHUVFTGCOPTQTO", "GOLANG"));
