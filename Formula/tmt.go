package main

// Tmt - Test Management Tool
// Homepage: https://tmt.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installTmt() {
	// Método 1: Descargar y extraer .tar.gz
	tmt_tar_url := "https://files.pythonhosted.org/packages/24/b7/7e7b018d011e2c045b47a9ee725244a3a68ba242a72ddc2e4e94cc1ef21e/tmt-1.35.0.tar.gz"
	tmt_cmd_tar := exec.Command("curl", "-L", tmt_tar_url, "-o", "package.tar.gz")
	err := tmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmt_zip_url := "https://files.pythonhosted.org/packages/24/b7/7e7b018d011e2c045b47a9ee725244a3a68ba242a72ddc2e4e94cc1ef21e/tmt-1.35.0.zip"
	tmt_cmd_zip := exec.Command("curl", "-L", tmt_zip_url, "-o", "package.zip")
	err = tmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmt_bin_url := "https://files.pythonhosted.org/packages/24/b7/7e7b018d011e2c045b47a9ee725244a3a68ba242a72ddc2e4e94cc1ef21e/tmt-1.35.0.bin"
	tmt_cmd_bin := exec.Command("curl", "-L", tmt_bin_url, "-o", "binary.bin")
	err = tmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmt_src_url := "https://files.pythonhosted.org/packages/24/b7/7e7b018d011e2c045b47a9ee725244a3a68ba242a72ddc2e4e94cc1ef21e/tmt-1.35.0.src.tar.gz"
	tmt_cmd_src := exec.Command("curl", "-L", tmt_src_url, "-o", "source.tar.gz")
	err = tmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmt_cmd_direct := exec.Command("./binary")
	err = tmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: beakerlib")
	exec.Command("latte", "install", "beakerlib").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
