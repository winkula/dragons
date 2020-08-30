<template>
  <div class="wrapper">
    <header>
      <Menu @new="newGame" @difficulty="changeDifficulty" @help="showHelp" @solve="solve"></Menu>
    </header>
    <main>
      <Grid :grid="game.puzzle" @filled="fillCell" :stataus="status" interactive></Grid>
    </main>
    <footer>
      <CellSelect v-model="fillType"></CellSelect>
      <div class="copyright">{{ copyright }}</div>
    </footer>
    <Help v-model="helpVisible"></Help>
    <Overlay :state="status" @close="overlayClosed"></Overlay>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

import Menu from "./Menu.vue";
import Grid from "./Grid.vue";
import CellSelect from "./CellSelect.vue";
import Help from "./Help.vue";
import Overlay from "./Overlay.vue";

import {
  createGame,
  Cell,
  CellDefinition,
  CellType,
  getCellType,
} from "../logic";
import { GameStatus } from "../logic/game";

export default Vue.extend({
  components: {
    Menu,
    Grid,
    CellSelect,
    Help,
    Overlay,
  },
  data() {
    return {
      game: null,
      status: null,
      isSolved: false,
      isValid: true,
      fillType: getCellType(CellType.Empty).value,
      helpVisible: false,
      copyright: "© 2020 Mathias Winkler",
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

      if (cell.isDefined) {
        // set back to undeinef
        cell.value = getCellType(CellType.Undefined).value;
      } else {
        // set value of cell
        cell.value = this.fillType;
      }

      // validate game
      this.status = this.game.status;
    },
    newGame() {
      this.game = createGame();
      this.status = "unsolved";
    },
    solve() {
      this.status = "solved";
    },
    changeDifficulty() {},
    showHelp() {
      this.helpVisible = !this.helpVisible;
    },
    overlayClosed() {
      this.status = "unsolved";
    },
  },
  created() {
    this.newGame();
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
    padding: 1rem 0;
  }

  & > main {
    padding: 1rem 0;
    background: lighten($color-background, 10%);
  }

  & > footer {
    padding: 1rem 0;
  }
}

html {
  background-color: $color-background;
  color: $color-font;
  user-select: none;
}

.copyright {
  margin-top: 1rem;
  text-align: center;
  color: $color-copyright;
}
</style>