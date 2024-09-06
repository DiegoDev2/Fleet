package main

// ZeldaRothSe - Zelda Return of the Hylian SE
// Homepage: https://www.solarus-games.org/en/games/the-legend-of-zelda-return-of-the-hylian-se

import (
	"fmt"
	
	"os/exec"
)

func installZeldaRothSe() {
	// Método 1: Descargar y extraer .tar.gz
	zeldarothse_tar_url := "https://gitlab.com/solarus-games/zelda-roth-se/-/archive/v1.2.1/zelda-roth-se-v1.2.1.tar.bz2"
	zeldarothse_cmd_tar := exec.Command("curl", "-L", zeldarothse_tar_url, "-o", "package.tar.gz")
	err := zeldarothse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zeldarothse_zip_url := "https://gitlab.com/solarus-games/zelda-roth-se/-/archive/v1.2.1/zelda-roth-se-v1.2.1.tar.bz2"
	zeldarothse_cmd_zip := exec.Command("curl", "-L", zeldarothse_zip_url, "-o", "package.zip")
	err = zeldarothse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zeldarothse_bin_url := "https://gitlab.com/solarus-games/zelda-roth-se/-/archive/v1.2.1/zelda-roth-se-v1.2.1.tar.bz2"
	zeldarothse_cmd_bin := exec.Command("curl", "-L", zeldarothse_bin_url, "-o", "binary.bin")
	err = zeldarothse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zeldarothse_src_url := "https://gitlab.com/solarus-games/zelda-roth-se/-/archive/v1.2.1/zelda-roth-se-v1.2.1.tar.bz2"
	zeldarothse_cmd_src := exec.Command("curl", "-L", zeldarothse_src_url, "-o", "source.tar.gz")
	err = zeldarothse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zeldarothse_cmd_direct := exec.Command("./binary")
	err = zeldarothse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: solarus")
exec.Command("latte", "install", "solarus")
}
