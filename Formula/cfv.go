package main

// Cfv - Test and create various files (e.g., .sfv, .csv, .crc., .torrent)
// Homepage: https://github.com/cfv-project/cfv

import (
	"fmt"
	
	"os/exec"
)

func installCfv() {
	// Método 1: Descargar y extraer .tar.gz
	cfv_tar_url := "https://files.pythonhosted.org/packages/29/ca/91cca3d1799d0e74b672e30c41f82a8135fe8d5baf7e6a8af2fdea282449/cfv-3.1.0.tar.gz"
	cfv_cmd_tar := exec.Command("curl", "-L", cfv_tar_url, "-o", "package.tar.gz")
	err := cfv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cfv_zip_url := "https://files.pythonhosted.org/packages/29/ca/91cca3d1799d0e74b672e30c41f82a8135fe8d5baf7e6a8af2fdea282449/cfv-3.1.0.zip"
	cfv_cmd_zip := exec.Command("curl", "-L", cfv_zip_url, "-o", "package.zip")
	err = cfv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cfv_bin_url := "https://files.pythonhosted.org/packages/29/ca/91cca3d1799d0e74b672e30c41f82a8135fe8d5baf7e6a8af2fdea282449/cfv-3.1.0.bin"
	cfv_cmd_bin := exec.Command("curl", "-L", cfv_bin_url, "-o", "binary.bin")
	err = cfv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cfv_src_url := "https://files.pythonhosted.org/packages/29/ca/91cca3d1799d0e74b672e30c41f82a8135fe8d5baf7e6a8af2fdea282449/cfv-3.1.0.src.tar.gz"
	cfv_cmd_src := exec.Command("curl", "-L", cfv_src_url, "-o", "source.tar.gz")
	err = cfv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cfv_cmd_direct := exec.Command("./binary")
	err = cfv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
