package main

// Chkbit - Check your files for data corruption
// Homepage: https://github.com/laktak/chkbit

import (
	"fmt"
	
	"os/exec"
)

func installChkbit() {
	// Método 1: Descargar y extraer .tar.gz
	chkbit_tar_url := "https://github.com/laktak/chkbit/archive/refs/tags/v5.2.0.tar.gz"
	chkbit_cmd_tar := exec.Command("curl", "-L", chkbit_tar_url, "-o", "package.tar.gz")
	err := chkbit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chkbit_zip_url := "https://github.com/laktak/chkbit/archive/refs/tags/v5.2.0.zip"
	chkbit_cmd_zip := exec.Command("curl", "-L", chkbit_zip_url, "-o", "package.zip")
	err = chkbit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chkbit_bin_url := "https://github.com/laktak/chkbit/archive/refs/tags/v5.2.0.bin"
	chkbit_cmd_bin := exec.Command("curl", "-L", chkbit_bin_url, "-o", "binary.bin")
	err = chkbit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chkbit_src_url := "https://github.com/laktak/chkbit/archive/refs/tags/v5.2.0.src.tar.gz"
	chkbit_cmd_src := exec.Command("curl", "-L", chkbit_src_url, "-o", "source.tar.gz")
	err = chkbit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chkbit_cmd_direct := exec.Command("./binary")
	err = chkbit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
