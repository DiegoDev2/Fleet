package main

// Stone - TCP/IP packet repeater in the application layer
// Homepage: https://www.gcd.org/sengoku/stone/

import (
	"fmt"
	
	"os/exec"
)

func installStone() {
	// Método 1: Descargar y extraer .tar.gz
	stone_tar_url := "https://www.gcd.org/sengoku/stone/stone-2.4.tar.gz"
	stone_cmd_tar := exec.Command("curl", "-L", stone_tar_url, "-o", "package.tar.gz")
	err := stone_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stone_zip_url := "https://www.gcd.org/sengoku/stone/stone-2.4.zip"
	stone_cmd_zip := exec.Command("curl", "-L", stone_zip_url, "-o", "package.zip")
	err = stone_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stone_bin_url := "https://www.gcd.org/sengoku/stone/stone-2.4.bin"
	stone_cmd_bin := exec.Command("curl", "-L", stone_bin_url, "-o", "binary.bin")
	err = stone_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stone_src_url := "https://www.gcd.org/sengoku/stone/stone-2.4.src.tar.gz"
	stone_cmd_src := exec.Command("curl", "-L", stone_src_url, "-o", "source.tar.gz")
	err = stone_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stone_cmd_direct := exec.Command("./binary")
	err = stone_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
