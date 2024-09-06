package main

// Sceptre - Build better AWS infrastructure
// Homepage: https://docs.sceptre-project.org/

import (
	"fmt"
	
	"os/exec"
)

func installSceptre() {
	// Método 1: Descargar y extraer .tar.gz
	sceptre_tar_url := "https://files.pythonhosted.org/packages/61/24/f82727943652a4f1be61bc2e08c1af2c562fe35343168bbe9c794511af9e/sceptre-4.5.2.tar.gz"
	sceptre_cmd_tar := exec.Command("curl", "-L", sceptre_tar_url, "-o", "package.tar.gz")
	err := sceptre_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sceptre_zip_url := "https://files.pythonhosted.org/packages/61/24/f82727943652a4f1be61bc2e08c1af2c562fe35343168bbe9c794511af9e/sceptre-4.5.2.zip"
	sceptre_cmd_zip := exec.Command("curl", "-L", sceptre_zip_url, "-o", "package.zip")
	err = sceptre_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sceptre_bin_url := "https://files.pythonhosted.org/packages/61/24/f82727943652a4f1be61bc2e08c1af2c562fe35343168bbe9c794511af9e/sceptre-4.5.2.bin"
	sceptre_cmd_bin := exec.Command("curl", "-L", sceptre_bin_url, "-o", "binary.bin")
	err = sceptre_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sceptre_src_url := "https://files.pythonhosted.org/packages/61/24/f82727943652a4f1be61bc2e08c1af2c562fe35343168bbe9c794511af9e/sceptre-4.5.2.src.tar.gz"
	sceptre_cmd_src := exec.Command("curl", "-L", sceptre_src_url, "-o", "source.tar.gz")
	err = sceptre_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sceptre_cmd_direct := exec.Command("./binary")
	err = sceptre_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
