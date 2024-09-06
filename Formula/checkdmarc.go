package main

// Checkdmarc - Command-line parser for SPF and DMARC DNS records
// Homepage: https://domainaware.github.io/checkdmarc/

import (
	"fmt"
	
	"os/exec"
)

func installCheckdmarc() {
	// Método 1: Descargar y extraer .tar.gz
	checkdmarc_tar_url := "https://files.pythonhosted.org/packages/b4/f3/0e4108ec0939ed21524e8fb2638725cca48ae750a837b09d3f6695092419/checkdmarc-5.5.0.tar.gz"
	checkdmarc_cmd_tar := exec.Command("curl", "-L", checkdmarc_tar_url, "-o", "package.tar.gz")
	err := checkdmarc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	checkdmarc_zip_url := "https://files.pythonhosted.org/packages/b4/f3/0e4108ec0939ed21524e8fb2638725cca48ae750a837b09d3f6695092419/checkdmarc-5.5.0.zip"
	checkdmarc_cmd_zip := exec.Command("curl", "-L", checkdmarc_zip_url, "-o", "package.zip")
	err = checkdmarc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	checkdmarc_bin_url := "https://files.pythonhosted.org/packages/b4/f3/0e4108ec0939ed21524e8fb2638725cca48ae750a837b09d3f6695092419/checkdmarc-5.5.0.bin"
	checkdmarc_cmd_bin := exec.Command("curl", "-L", checkdmarc_bin_url, "-o", "binary.bin")
	err = checkdmarc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	checkdmarc_src_url := "https://files.pythonhosted.org/packages/b4/f3/0e4108ec0939ed21524e8fb2638725cca48ae750a837b09d3f6695092419/checkdmarc-5.5.0.src.tar.gz"
	checkdmarc_cmd_src := exec.Command("curl", "-L", checkdmarc_src_url, "-o", "source.tar.gz")
	err = checkdmarc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	checkdmarc_cmd_direct := exec.Command("./binary")
	err = checkdmarc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
