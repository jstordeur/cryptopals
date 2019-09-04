import base64

def xor(input, key):
    """Returns a bytearray of input xored with key. If key is shorter than input, loop over key.
    """
    res = bytearray()
    index = 0
    assert len(key) > 0, "key can't be empty"
    for i in input:
        res.append(i^key[index])
        index = (index+1)%len(key)
    return res

def set1():
    def challenge1():
        input = ('49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d75736872'
                 '6f6f6d')
        expected = bytearray('SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t', 'utf-8')
        res = base64.b64encode(bytearray.fromhex(input))
        assert res == expected, '\nexpected: {}\nres: {}\n'.format(expected, res)

    def challenge2():
        input = bytearray.fromhex('1c0111001f010100061a024b53535009181c')
        key = bytearray.fromhex('686974207468652062756c6c277320657965')
        res = xor(input, key)
        expected = bytearray.fromhex('746865206b696420646f6e277420706c6179')
        assert res == expected, '\nexpected: {}\nres: {}\n'.format(expected, res)


    challenge1()
    challenge2()

def main():
    set1()


if __name__=="__main__":
    main()
