package main

// Snakemake - Pythonic workflow system
// Homepage: https://snakemake.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installSnakemake() {
	// Método 1: Descargar y extraer .tar.gz
	snakemake_tar_url := "https://files.pythonhosted.org/packages/a9/cc/c15d378c3e956f10df82f7c3f58206548ecfafa705a54c7bee5d1a8c2935/snakemake-8.19.3.tar.gz"
	snakemake_cmd_tar := exec.Command("curl", "-L", snakemake_tar_url, "-o", "package.tar.gz")
	err := snakemake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snakemake_zip_url := "https://files.pythonhosted.org/packages/a9/cc/c15d378c3e956f10df82f7c3f58206548ecfafa705a54c7bee5d1a8c2935/snakemake-8.19.3.zip"
	snakemake_cmd_zip := exec.Command("curl", "-L", snakemake_zip_url, "-o", "package.zip")
	err = snakemake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snakemake_bin_url := "https://files.pythonhosted.org/packages/a9/cc/c15d378c3e956f10df82f7c3f58206548ecfafa705a54c7bee5d1a8c2935/snakemake-8.19.3.bin"
	snakemake_cmd_bin := exec.Command("curl", "-L", snakemake_bin_url, "-o", "binary.bin")
	err = snakemake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snakemake_src_url := "https://files.pythonhosted.org/packages/a9/cc/c15d378c3e956f10df82f7c3f58206548ecfafa705a54c7bee5d1a8c2935/snakemake-8.19.3.src.tar.gz"
	snakemake_cmd_src := exec.Command("curl", "-L", snakemake_src_url, "-o", "source.tar.gz")
	err = snakemake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snakemake_cmd_direct := exec.Command("./binary")
	err = snakemake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: cbc")
exec.Command("latte", "install", "cbc")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
