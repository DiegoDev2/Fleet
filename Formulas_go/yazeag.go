package main

// YazeAg - Yet Another Z80 Emulator (by AG)
// Homepage: https://www.mathematik.uni-ulm.de/users/ag/yaze-ag/

import (
	"fmt"
	
	"os/exec"
)

func installYazeAg() {
	// Método 1: Descargar y extraer .tar.gz
	yazeag_tar_url := "https://www.mathematik.uni-ulm.de/users/ag/yaze-ag/devel/yaze-ag-2.51.3.tar.gz"
	yazeag_cmd_tar := exec.Command("curl", "-L", yazeag_tar_url, "-o", "package.tar.gz")
	err := yazeag_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yazeag_zip_url := "https://www.mathematik.uni-ulm.de/users/ag/yaze-ag/devel/yaze-ag-2.51.3.zip"
	yazeag_cmd_zip := exec.Command("curl", "-L", yazeag_zip_url, "-o", "package.zip")
	err = yazeag_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yazeag_bin_url := "https://www.mathematik.uni-ulm.de/users/ag/yaze-ag/devel/yaze-ag-2.51.3.bin"
	yazeag_cmd_bin := exec.Command("curl", "-L", yazeag_bin_url, "-o", "binary.bin")
	err = yazeag_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yazeag_src_url := "https://www.mathematik.uni-ulm.de/users/ag/yaze-ag/devel/yaze-ag-2.51.3.src.tar.gz"
	yazeag_cmd_src := exec.Command("curl", "-L", yazeag_src_url, "-o", "source.tar.gz")
	err = yazeag_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yazeag_cmd_direct := exec.Command("./binary")
	err = yazeag_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
