package main

// TranslateToolkit - Toolkit for localization engineers
// Homepage: https://toolkit.translatehouse.org/

import (
	"fmt"
	
	"os/exec"
)

func installTranslateToolkit() {
	// Método 1: Descargar y extraer .tar.gz
	translatetoolkit_tar_url := "https://files.pythonhosted.org/packages/0a/73/21e21d5648e7d599c476816695fe31e8b059c2f2e0b99804cd5ca9366994/translate_toolkit-3.13.3.tar.gz"
	translatetoolkit_cmd_tar := exec.Command("curl", "-L", translatetoolkit_tar_url, "-o", "package.tar.gz")
	err := translatetoolkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	translatetoolkit_zip_url := "https://files.pythonhosted.org/packages/0a/73/21e21d5648e7d599c476816695fe31e8b059c2f2e0b99804cd5ca9366994/translate_toolkit-3.13.3.zip"
	translatetoolkit_cmd_zip := exec.Command("curl", "-L", translatetoolkit_zip_url, "-o", "package.zip")
	err = translatetoolkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	translatetoolkit_bin_url := "https://files.pythonhosted.org/packages/0a/73/21e21d5648e7d599c476816695fe31e8b059c2f2e0b99804cd5ca9366994/translate_toolkit-3.13.3.bin"
	translatetoolkit_cmd_bin := exec.Command("curl", "-L", translatetoolkit_bin_url, "-o", "binary.bin")
	err = translatetoolkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	translatetoolkit_src_url := "https://files.pythonhosted.org/packages/0a/73/21e21d5648e7d599c476816695fe31e8b059c2f2e0b99804cd5ca9366994/translate_toolkit-3.13.3.src.tar.gz"
	translatetoolkit_cmd_src := exec.Command("curl", "-L", translatetoolkit_src_url, "-o", "source.tar.gz")
	err = translatetoolkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	translatetoolkit_cmd_direct := exec.Command("./binary")
	err = translatetoolkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
