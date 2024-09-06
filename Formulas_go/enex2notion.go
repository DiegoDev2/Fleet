package main

// Enex2notion - Import Evernote ENEX files to Notion
// Homepage: https://github.com/vzhd1701/enex2notion

import (
	"fmt"
	
	"os/exec"
)

func installEnex2notion() {
	// Método 1: Descargar y extraer .tar.gz
	enex2notion_tar_url := "https://files.pythonhosted.org/packages/de/5c/c0ce22d810226345411b03177f9b43c35b82c3a671d5d73f56fc43b0858e/enex2notion-0.3.1.tar.gz"
	enex2notion_cmd_tar := exec.Command("curl", "-L", enex2notion_tar_url, "-o", "package.tar.gz")
	err := enex2notion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	enex2notion_zip_url := "https://files.pythonhosted.org/packages/de/5c/c0ce22d810226345411b03177f9b43c35b82c3a671d5d73f56fc43b0858e/enex2notion-0.3.1.zip"
	enex2notion_cmd_zip := exec.Command("curl", "-L", enex2notion_zip_url, "-o", "package.zip")
	err = enex2notion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	enex2notion_bin_url := "https://files.pythonhosted.org/packages/de/5c/c0ce22d810226345411b03177f9b43c35b82c3a671d5d73f56fc43b0858e/enex2notion-0.3.1.bin"
	enex2notion_cmd_bin := exec.Command("curl", "-L", enex2notion_bin_url, "-o", "binary.bin")
	err = enex2notion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	enex2notion_src_url := "https://files.pythonhosted.org/packages/de/5c/c0ce22d810226345411b03177f9b43c35b82c3a671d5d73f56fc43b0858e/enex2notion-0.3.1.src.tar.gz"
	enex2notion_cmd_src := exec.Command("curl", "-L", enex2notion_src_url, "-o", "source.tar.gz")
	err = enex2notion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	enex2notion_cmd_direct := exec.Command("./binary")
	err = enex2notion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: pymupdf")
exec.Command("latte", "install", "pymupdf")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
