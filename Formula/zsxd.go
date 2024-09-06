package main

// Zsxd - Zelda Mystery of Solarus XD
// Homepage: https://www.solarus-games.org/games/the-legend-of-zelda-mystery-of-solarus-xd/

import (
	"fmt"
	
	"os/exec"
)

func installZsxd() {
	// Método 1: Descargar y extraer .tar.gz
	zsxd_tar_url := "https://gitlab.com/solarus-games/games/zsxd/-/archive/v1.12.2/zsxd-v1.12.2.tar.bz2"
	zsxd_cmd_tar := exec.Command("curl", "-L", zsxd_tar_url, "-o", "package.tar.gz")
	err := zsxd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zsxd_zip_url := "https://gitlab.com/solarus-games/games/zsxd/-/archive/v1.12.2/zsxd-v1.12.2.tar.bz2"
	zsxd_cmd_zip := exec.Command("curl", "-L", zsxd_zip_url, "-o", "package.zip")
	err = zsxd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zsxd_bin_url := "https://gitlab.com/solarus-games/games/zsxd/-/archive/v1.12.2/zsxd-v1.12.2.tar.bz2"
	zsxd_cmd_bin := exec.Command("curl", "-L", zsxd_bin_url, "-o", "binary.bin")
	err = zsxd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zsxd_src_url := "https://gitlab.com/solarus-games/games/zsxd/-/archive/v1.12.2/zsxd-v1.12.2.tar.bz2"
	zsxd_cmd_src := exec.Command("curl", "-L", zsxd_src_url, "-o", "source.tar.gz")
	err = zsxd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zsxd_cmd_direct := exec.Command("./binary")
	err = zsxd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: solarus")
	exec.Command("latte", "install", "solarus").Run()
}
