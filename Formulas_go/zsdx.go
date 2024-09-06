package main

// Zsdx - Zelda Mystery of Solarus DX
// Homepage: https://www.solarus-games.org/games/the-legend-of-zelda-mystery-of-solarus-dx

import (
	"fmt"
	
	"os/exec"
)

func installZsdx() {
	// Método 1: Descargar y extraer .tar.gz
	zsdx_tar_url := "https://gitlab.com/solarus-games/games/zsdx/-/archive/v1.12.3/zsdx-v1.12.3.tar.bz2"
	zsdx_cmd_tar := exec.Command("curl", "-L", zsdx_tar_url, "-o", "package.tar.gz")
	err := zsdx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zsdx_zip_url := "https://gitlab.com/solarus-games/games/zsdx/-/archive/v1.12.3/zsdx-v1.12.3.tar.bz2"
	zsdx_cmd_zip := exec.Command("curl", "-L", zsdx_zip_url, "-o", "package.zip")
	err = zsdx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zsdx_bin_url := "https://gitlab.com/solarus-games/games/zsdx/-/archive/v1.12.3/zsdx-v1.12.3.tar.bz2"
	zsdx_cmd_bin := exec.Command("curl", "-L", zsdx_bin_url, "-o", "binary.bin")
	err = zsdx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zsdx_src_url := "https://gitlab.com/solarus-games/games/zsdx/-/archive/v1.12.3/zsdx-v1.12.3.tar.bz2"
	zsdx_cmd_src := exec.Command("curl", "-L", zsdx_src_url, "-o", "source.tar.gz")
	err = zsdx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zsdx_cmd_direct := exec.Command("./binary")
	err = zsdx_cmd_direct.Run()
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
