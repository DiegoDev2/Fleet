package main

// Rbtools - CLI and API for working with code and document reviews on Review Board
// Homepage: https://www.reviewboard.org/downloads/rbtools/

import (
	"fmt"
	
	"os/exec"
)

func installRbtools() {
	// Método 1: Descargar y extraer .tar.gz
	rbtools_tar_url := "https://files.pythonhosted.org/packages/47/20/112a54002d2023b249906508feefeebcce7d73a25618f3f1f76c7673734d/RBTools-5.0.tar.gz"
	rbtools_cmd_tar := exec.Command("curl", "-L", rbtools_tar_url, "-o", "package.tar.gz")
	err := rbtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbtools_zip_url := "https://files.pythonhosted.org/packages/47/20/112a54002d2023b249906508feefeebcce7d73a25618f3f1f76c7673734d/RBTools-5.0.zip"
	rbtools_cmd_zip := exec.Command("curl", "-L", rbtools_zip_url, "-o", "package.zip")
	err = rbtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbtools_bin_url := "https://files.pythonhosted.org/packages/47/20/112a54002d2023b249906508feefeebcce7d73a25618f3f1f76c7673734d/RBTools-5.0.bin"
	rbtools_cmd_bin := exec.Command("curl", "-L", rbtools_bin_url, "-o", "binary.bin")
	err = rbtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbtools_src_url := "https://files.pythonhosted.org/packages/47/20/112a54002d2023b249906508feefeebcce7d73a25618f3f1f76c7673734d/RBTools-5.0.src.tar.gz"
	rbtools_cmd_src := exec.Command("curl", "-L", rbtools_src_url, "-o", "source.tar.gz")
	err = rbtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbtools_cmd_direct := exec.Command("./binary")
	err = rbtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
