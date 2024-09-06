package main

// GiDocgen - Documentation tool for GObject-based libraries
// Homepage: https://gnome.pages.gitlab.gnome.org/gi-docgen/

import (
	"fmt"
	
	"os/exec"
)

func installGiDocgen() {
	// Método 1: Descargar y extraer .tar.gz
	gidocgen_tar_url := "https://files.pythonhosted.org/packages/d1/86/d17f162d174b6340031fc96474405f13d50ceda4b6bf6588593cf31eb84b/gi_docgen-2024.1.tar.gz"
	gidocgen_cmd_tar := exec.Command("curl", "-L", gidocgen_tar_url, "-o", "package.tar.gz")
	err := gidocgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gidocgen_zip_url := "https://files.pythonhosted.org/packages/d1/86/d17f162d174b6340031fc96474405f13d50ceda4b6bf6588593cf31eb84b/gi_docgen-2024.1.zip"
	gidocgen_cmd_zip := exec.Command("curl", "-L", gidocgen_zip_url, "-o", "package.zip")
	err = gidocgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gidocgen_bin_url := "https://files.pythonhosted.org/packages/d1/86/d17f162d174b6340031fc96474405f13d50ceda4b6bf6588593cf31eb84b/gi_docgen-2024.1.bin"
	gidocgen_cmd_bin := exec.Command("curl", "-L", gidocgen_bin_url, "-o", "binary.bin")
	err = gidocgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gidocgen_src_url := "https://files.pythonhosted.org/packages/d1/86/d17f162d174b6340031fc96474405f13d50ceda4b6bf6588593cf31eb84b/gi_docgen-2024.1.src.tar.gz"
	gidocgen_cmd_src := exec.Command("curl", "-L", gidocgen_src_url, "-o", "source.tar.gz")
	err = gidocgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gidocgen_cmd_direct := exec.Command("./binary")
	err = gidocgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
