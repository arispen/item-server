import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>item-server client</title>
        <meta name="description" content="item-server client" />
        {/* <link rel="icon" href="/favicon.ico" /> */}
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>Welcome to item-server</h1>
        <div>
          <Image
            src="/jester-hat.png"
            width={256}
            height={256}
          />
        </div>
      </main>

      <footer className={styles.footer}>powered by magic</footer>
    </div>
  );
}
