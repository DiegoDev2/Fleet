package main

// RichCli - Command-line toolbox for fancy output in the terminal
// Homepage: https://github.com/textualize/rich-cli

import (
	"fmt"
	
	"os/exec"
)

func installRichCli() {
	// Método 1: Descargar y extraer .tar.gz
	richcli_tar_url := "https://files.pythonhosted.org/packages/ca/55/e35962573948a148a4f63416d95d25fe75feb06d9ae2f9bb35adc416f894/rich-cli-1.8.0.tar.gz"
	richcli_cmd_tar := exec.Command("curl", "-L", richcli_tar_url, "-o", "package.tar.gz")
	err := richcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	richcli_zip_url := "https://files.pythonhosted.org/packages/ca/55/e35962573948a148a4f63416d95d25fe75feb06d9ae2f9bb35adc416f894/rich-cli-1.8.0.zip"
	richcli_cmd_zip := exec.Command("curl", "-L", richcli_zip_url, "-o", "package.zip")
	err = richcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	richcli_bin_url := "https://files.pythonhosted.org/packages/ca/55/e35962573948a148a4f63416d95d25fe75feb06d9ae2f9bb35adc416f894/rich-cli-1.8.0.bin"
	richcli_cmd_bin := exec.Command("curl", "-L", richcli_bin_url, "-o", "binary.bin")
	err = richcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	richcli_src_url := "https://files.pythonhosted.org/packages/ca/55/e35962573948a148a4f63416d95d25fe75feb06d9ae2f9bb35adc416f894/rich-cli-1.8.0.src.tar.gz"
	richcli_cmd_src := exec.Command("curl", "-L", richcli_src_url, "-o", "source.tar.gz")
	err = richcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	richcli_cmd_direct := exec.Command("./binary")
	err = richcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
