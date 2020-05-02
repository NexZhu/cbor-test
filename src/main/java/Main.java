import co.nstant.in.cbor.CborBuilder;
import co.nstant.in.cbor.CborEncoder;
import co.nstant.in.cbor.CborException;

import java.io.ByteArrayOutputStream;
import java.io.FileOutputStream;
import java.io.IOException;

class Transaction {
  public int type;
  public long nonce;
  public byte[] src_addr;
  public byte[] dst_addr;
  public byte[] data;
  public byte[] block_hash;
}

public class Main {
  public static void main(String[] args) throws CborException, IOException {

    var tx = new Transaction();
    tx.type = 1;
    tx.nonce = 52;
    tx.src_addr = new byte[]{0x1, 0x2, 0x3, 0x4};
    tx.dst_addr = new byte[]{0x5, 0x6, 0x7, 0x8};
    tx.data = new byte[]{0x1, 0x2, 0x3, 0x4};
    tx.block_hash = new byte[]{0x5, 0x6, 0x7, 0x8};

    var baos = new ByteArrayOutputStream();
    new CborEncoder(baos).encode(new CborBuilder()
        .addMap()
        .put("type", tx.type)
        .put("nonce", tx.nonce)
        .put("src_addr", tx.src_addr)
        .put("dst_addr", tx.dst_addr)
        .put("data", tx.data)
        .put("block_hash", tx.block_hash)
        .end()
        .build());
//    byte[] bytes = baos.toByteArray();
    var fos = new FileOutputStream("java.txt");
    baos.writeTo(fos);
    baos.close();
    fos.close();
//    var tx2 = mapper.readValue(cborData, Transaction.class);
//    System.out.println(tx2.nonce);
//    System.out.println(tx2.src_addr[3]);
  }
}
