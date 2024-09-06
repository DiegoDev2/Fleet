package main

// Weasyprint - Convert HTML to PDF
// Homepage: https://www.courtbouillon.org/weasyprint

import (
	"fmt"
	
	"os/exec"
)

func installWeasyprint() {
	// Método 1: Descargar y extraer .tar.gz
	weasyprint_tar_url := "https://files.pythonhosted.org/packages/fd/22/2d76310cd2ea5bbf03c691a08d48626f49853b7261a51bbdc0f834d746ca/weasyprint-62.3.tar.gz"
	weasyprint_cmd_tar := exec.Command("curl", "-L", weasyprint_tar_url, "-o", "package.tar.gz")
	err := weasyprint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	weasyprint_zip_url := "https://files.pythonhosted.org/packages/fd/22/2d76310cd2ea5bbf03c691a08d48626f49853b7261a51bbdc0f834d746ca/weasyprint-62.3.zip"
	weasyprint_cmd_zip := exec.Command("curl", "-L", weasyprint_zip_url, "-o", "package.zip")
	err = weasyprint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	weasyprint_bin_url := "https://files.pythonhosted.org/packages/fd/22/2d76310cd2ea5bbf03c691a08d48626f49853b7261a51bbdc0f834d746ca/weasyprint-62.3.bin"
	weasyprint_cmd_bin := exec.Command("curl", "-L", weasyprint_bin_url, "-o", "binary.bin")
	err = weasyprint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	weasyprint_src_url := "https://files.pythonhosted.org/packages/fd/22/2d76310cd2ea5bbf03c691a08d48626f49853b7261a51bbdc0f834d746ca/weasyprint-62.3.src.tar.gz"
	weasyprint_cmd_src := exec.Command("curl", "-L", weasyprint_src_url, "-o", "source.tar.gz")
	err = weasyprint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	weasyprint_cmd_direct := exec.Command("./binary")
	err = weasyprint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
