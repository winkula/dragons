<template>
  <div class="wrapper" v-cloak>
    <header>
      <Menu @new="openStartDialog" @help="showHelp" @solve="solve"></Menu>
    </header>
    <main>
      <Grid :grid="game.puzzle" @filled="fillCell" :stataus="status" interactive></Grid>
    </main>
    <footer>
      <CellSelect v-model="fillType"></CellSelect>
    </footer>
    <Help v-model="helpVisible"></Help>
    <Overlay :state="status" @close="overlayClosed"></Overlay>
    <StartDialog v-model="startDialogVisible" @start="start"></StartDialog>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

import Menu from "./Menu.vue";
import Grid from "./Grid.vue";
import CellSelect from "./CellSelect.vue";
import Help from "./Help.vue";
import Overlay from "./Overlay.vue";
import StartDialog from "./StartDialog.vue";

import {
  createGame,
  emptyGame,
  Cell,
  CellDefinition,
  CellType,
  getCellType,
} from "../logic";
import { Difficulty, GameStatus } from "../logic/game";
import { playClick, playMusic, playError, playWin } from "../sfx";

export default Vue.extend({
  components: {
    Menu,
    Grid,
    CellSelect,
    Help,
    Overlay,
    StartDialog,
  },
  data() {
    return {
      game: emptyGame,
      size: 8,
      difficulty: "easy",
      status: "unsolved",
      isSolved: false,
      isValid: true,
      fillType: getCellType(CellType.Air).value,
      helpVisible: false,
      startDialogVisible: true,
    };
  },
  methods: {
    fillCell(cell: Cell) {
      if (cell.given) {
        // cannot change a cell that is given
        return;
      }

      if (this.fillType == null) {
        // cannot set a cell, when no fill type is selected
        return;
      }

      if (cell.value === this.fillType) {
        // set back to undeinef
        cell.value = getCellType(CellType.Undefined).value;
      } else {
        // set value of cell
        cell.value = this.fillType;
      }

      // validate game
      this.status = this.game.status;
      if (this.status === GameStatus.Invalid) {
        playError();
      } else if (this.status === GameStatus.Solved) {
        playWin();
      } else {
        playClick();
      }
    },
    newGame() {
      this.game = createGame(this.difficulty as Difficulty, this.size);
      this.status = "unsolved";
    },
    solve() {
      this.status = "solved";
    },
    showHelp() {
      this.helpVisible = !this.helpVisible;
    },
    overlayClosed() {
      this.status = "unsolved";
    },
    start(difficulty) {
      playMusic();
      this.difficulty = difficulty;
      this.startDialogVisible = false;
      this.newGame();
    },
    openStartDialog() {
      this.startDialogVisible = true;
    },
  },
});
</script>

<style lang="scss">
@import "./src/assets/styles/globals";

html,
body,
.wrapper {
  height: 100%;
  margin: 0;
  padding: 0;
}

.wrapper {
  display: flex;
  flex-direction: column;
  justify-content: center;

  & > header {
    padding: 2vmin 0;
  }

  & > main {
    padding: 2vmin 0;
    background: $color-background-inner;
  }

  & > footer {
    padding: 2vmin 0;
  }
}

html {
  color: $color-font;
  user-select: none;
}

body {
  background-color: $color-background;
  background: linear-gradient(
    $color-background-inner,
    $color-background,
    $color-background,
    $color-background-inner
  );
}

p.copyright {
  margin-top: 1rem;
  text-align: center;
  color: $color-font-copyright;
}
</style>