package main

// Tern - Software Bill of Materials (SBOM) tool
// Homepage: https://github.com/tern-tools/tern

import (
	"fmt"
	
	"os/exec"
)

func installTern() {
	// Método 1: Descargar y extraer .tar.gz
	tern_tar_url := "https://files.pythonhosted.org/packages/f8/4b/123b2ca469126b45e61853acf028fe1d466f4fe1d5e7afd1d4972c151b4d/tern-2.12.1.tar.gz"
	tern_cmd_tar := exec.Command("curl", "-L", tern_tar_url, "-o", "package.tar.gz")
	err := tern_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tern_zip_url := "https://files.pythonhosted.org/packages/f8/4b/123b2ca469126b45e61853acf028fe1d466f4fe1d5e7afd1d4972c151b4d/tern-2.12.1.zip"
	tern_cmd_zip := exec.Command("curl", "-L", tern_zip_url, "-o", "package.zip")
	err = tern_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tern_bin_url := "https://files.pythonhosted.org/packages/f8/4b/123b2ca469126b45e61853acf028fe1d466f4fe1d5e7afd1d4972c151b4d/tern-2.12.1.bin"
	tern_cmd_bin := exec.Command("curl", "-L", tern_bin_url, "-o", "binary.bin")
	err = tern_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tern_src_url := "https://files.pythonhosted.org/packages/f8/4b/123b2ca469126b45e61853acf028fe1d466f4fe1d5e7afd1d4972c151b4d/tern-2.12.1.src.tar.gz"
	tern_cmd_src := exec.Command("curl", "-L", tern_src_url, "-o", "source.tar.gz")
	err = tern_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tern_cmd_direct := exec.Command("./binary")
	err = tern_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: attr")
exec.Command("latte", "install", "attr")
}
