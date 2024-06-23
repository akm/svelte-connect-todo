import { initializeApp } from 'firebase/app';

// ui/src/lib/firebase/fbconfig.ts は Firebase のコンソールで
// プロジェクトの設定 > 全般 > マイアプリ > SDK の設定と構成 の Config を選択して
// 表示されるコードをコピーして作成してください。
import { firebaseConfig } from './fbconfig';

export const app = initializeApp(firebaseConfig);
