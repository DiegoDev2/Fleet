package main

// Rpl - Text replacement utility
// Homepage: https://github.com/rrthomas/rpl

import (
	"fmt"
	
	"os/exec"
)

func installRpl() {
	// Método 1: Descargar y extraer .tar.gz
	rpl_tar_url := "https://files.pythonhosted.org/packages/40/ad/840b679493c49e0c4368662e2ddd6296f9bac41e8ee992e0d43d144b4f35/rpl-1.15.7.tar.gz"
	rpl_cmd_tar := exec.Command("curl", "-L", rpl_tar_url, "-o", "package.tar.gz")
	err := rpl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rpl_zip_url := "https://files.pythonhosted.org/packages/40/ad/840b679493c49e0c4368662e2ddd6296f9bac41e8ee992e0d43d144b4f35/rpl-1.15.7.zip"
	rpl_cmd_zip := exec.Command("curl", "-L", rpl_zip_url, "-o", "package.zip")
	err = rpl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rpl_bin_url := "https://files.pythonhosted.org/packages/40/ad/840b679493c49e0c4368662e2ddd6296f9bac41e8ee992e0d43d144b4f35/rpl-1.15.7.bin"
	rpl_cmd_bin := exec.Command("curl", "-L", rpl_bin_url, "-o", "binary.bin")
	err = rpl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rpl_src_url := "https://files.pythonhosted.org/packages/40/ad/840b679493c49e0c4368662e2ddd6296f9bac41e8ee992e0d43d144b4f35/rpl-1.15.7.src.tar.gz"
	rpl_cmd_src := exec.Command("curl", "-L", rpl_src_url, "-o", "source.tar.gz")
	err = rpl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rpl_cmd_direct := exec.Command("./binary")
	err = rpl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
