package main

// Dtm - Cross-language distributed transaction manager
// Homepage: https://en.dtm.pub/

import (
	"fmt"
	
	"os/exec"
)

func installDtm() {
	// Método 1: Descargar y extraer .tar.gz
	dtm_tar_url := "https://github.com/dtm-labs/dtm/archive/refs/tags/v1.18.0.tar.gz"
	dtm_cmd_tar := exec.Command("curl", "-L", dtm_tar_url, "-o", "package.tar.gz")
	err := dtm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dtm_zip_url := "https://github.com/dtm-labs/dtm/archive/refs/tags/v1.18.0.zip"
	dtm_cmd_zip := exec.Command("curl", "-L", dtm_zip_url, "-o", "package.zip")
	err = dtm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dtm_bin_url := "https://github.com/dtm-labs/dtm/archive/refs/tags/v1.18.0.bin"
	dtm_cmd_bin := exec.Command("curl", "-L", dtm_bin_url, "-o", "binary.bin")
	err = dtm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dtm_src_url := "https://github.com/dtm-labs/dtm/archive/refs/tags/v1.18.0.src.tar.gz"
	dtm_cmd_src := exec.Command("curl", "-L", dtm_src_url, "-o", "source.tar.gz")
	err = dtm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dtm_cmd_direct := exec.Command("./binary")
	err = dtm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
