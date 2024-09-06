package main

// Rogue - Dungeon crawling video game
// Homepage: https://sourceforge.net/projects/roguelike/

import (
	"fmt"
	
	"os/exec"
)

func installRogue() {
	// Método 1: Descargar y extraer .tar.gz
	rogue_tar_url := "https://src.fedoraproject.org/repo/pkgs/rogue/rogue5.4.4-src.tar.gz/033288f46444b06814c81ea69d96e075/rogue5.4.4-src.tar.gz"
	rogue_cmd_tar := exec.Command("curl", "-L", rogue_tar_url, "-o", "package.tar.gz")
	err := rogue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rogue_zip_url := "https://src.fedoraproject.org/repo/pkgs/rogue/rogue5.4.4-src.zip/033288f46444b06814c81ea69d96e075/rogue5.4.4-src.zip"
	rogue_cmd_zip := exec.Command("curl", "-L", rogue_zip_url, "-o", "package.zip")
	err = rogue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rogue_bin_url := "https://src.fedoraproject.org/repo/pkgs/rogue/rogue5.4.4-src.bin/033288f46444b06814c81ea69d96e075/rogue5.4.4-src.bin"
	rogue_cmd_bin := exec.Command("curl", "-L", rogue_bin_url, "-o", "binary.bin")
	err = rogue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rogue_src_url := "https://src.fedoraproject.org/repo/pkgs/rogue/rogue5.4.4-src.src.tar.gz/033288f46444b06814c81ea69d96e075/rogue5.4.4-src.src.tar.gz"
	rogue_cmd_src := exec.Command("curl", "-L", rogue_src_url, "-o", "source.tar.gz")
	err = rogue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rogue_cmd_direct := exec.Command("./binary")
	err = rogue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
