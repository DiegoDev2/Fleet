package main

// Dungeon - Classic text adventure game
// Homepage: https://github.com/GOFAI/dungeon

import (
	"fmt"
	
	"os/exec"
)

func installDungeon() {
	// Método 1: Descargar y extraer .tar.gz
	dungeon_tar_url := "https://github.com/GOFAI/dungeon/archive/refs/tags/4.1.tar.gz"
	dungeon_cmd_tar := exec.Command("curl", "-L", dungeon_tar_url, "-o", "package.tar.gz")
	err := dungeon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dungeon_zip_url := "https://github.com/GOFAI/dungeon/archive/refs/tags/4.1.zip"
	dungeon_cmd_zip := exec.Command("curl", "-L", dungeon_zip_url, "-o", "package.zip")
	err = dungeon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dungeon_bin_url := "https://github.com/GOFAI/dungeon/archive/refs/tags/4.1.bin"
	dungeon_cmd_bin := exec.Command("curl", "-L", dungeon_bin_url, "-o", "binary.bin")
	err = dungeon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dungeon_src_url := "https://github.com/GOFAI/dungeon/archive/refs/tags/4.1.src.tar.gz"
	dungeon_cmd_src := exec.Command("curl", "-L", dungeon_src_url, "-o", "source.tar.gz")
	err = dungeon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dungeon_cmd_direct := exec.Command("./binary")
	err = dungeon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}
