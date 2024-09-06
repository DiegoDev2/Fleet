package main

// Datalad - Data distribution geared toward scientific datasets
// Homepage: https://www.datalad.org

import (
	"fmt"
	
	"os/exec"
)

func installDatalad() {
	// Método 1: Descargar y extraer .tar.gz
	datalad_tar_url := "https://files.pythonhosted.org/packages/17/67/ffd33d1011477b0f87975dc36aef3817a7d3a8932678ce959583d90f4ebb/datalad-1.1.3.tar.gz"
	datalad_cmd_tar := exec.Command("curl", "-L", datalad_tar_url, "-o", "package.tar.gz")
	err := datalad_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	datalad_zip_url := "https://files.pythonhosted.org/packages/17/67/ffd33d1011477b0f87975dc36aef3817a7d3a8932678ce959583d90f4ebb/datalad-1.1.3.zip"
	datalad_cmd_zip := exec.Command("curl", "-L", datalad_zip_url, "-o", "package.zip")
	err = datalad_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	datalad_bin_url := "https://files.pythonhosted.org/packages/17/67/ffd33d1011477b0f87975dc36aef3817a7d3a8932678ce959583d90f4ebb/datalad-1.1.3.bin"
	datalad_cmd_bin := exec.Command("curl", "-L", datalad_bin_url, "-o", "binary.bin")
	err = datalad_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	datalad_src_url := "https://files.pythonhosted.org/packages/17/67/ffd33d1011477b0f87975dc36aef3817a7d3a8932678ce959583d90f4ebb/datalad-1.1.3.src.tar.gz"
	datalad_cmd_src := exec.Command("curl", "-L", datalad_src_url, "-o", "source.tar.gz")
	err = datalad_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	datalad_cmd_direct := exec.Command("./binary")
	err = datalad_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: git-annex")
	exec.Command("latte", "install", "git-annex").Run()
	fmt.Println("Instalando dependencia: p7zip")
	exec.Command("latte", "install", "p7zip").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
