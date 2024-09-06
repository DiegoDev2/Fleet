package main

// OrganizeTool - File management automation tool
// Homepage: https://github.com/tfeldmann/organize

import (
	"fmt"
	
	"os/exec"
)

func installOrganizeTool() {
	// Método 1: Descargar y extraer .tar.gz
	organizetool_tar_url := "https://files.pythonhosted.org/packages/ef/45/6b36a81132cd91b35f6727826533e9f166070eace72f3a09d85e7829c515/organize_tool-3.2.5.tar.gz"
	organizetool_cmd_tar := exec.Command("curl", "-L", organizetool_tar_url, "-o", "package.tar.gz")
	err := organizetool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	organizetool_zip_url := "https://files.pythonhosted.org/packages/ef/45/6b36a81132cd91b35f6727826533e9f166070eace72f3a09d85e7829c515/organize_tool-3.2.5.zip"
	organizetool_cmd_zip := exec.Command("curl", "-L", organizetool_zip_url, "-o", "package.zip")
	err = organizetool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	organizetool_bin_url := "https://files.pythonhosted.org/packages/ef/45/6b36a81132cd91b35f6727826533e9f166070eace72f3a09d85e7829c515/organize_tool-3.2.5.bin"
	organizetool_cmd_bin := exec.Command("curl", "-L", organizetool_bin_url, "-o", "binary.bin")
	err = organizetool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	organizetool_src_url := "https://files.pythonhosted.org/packages/ef/45/6b36a81132cd91b35f6727826533e9f166070eace72f3a09d85e7829c515/organize_tool-3.2.5.src.tar.gz"
	organizetool_cmd_src := exec.Command("curl", "-L", organizetool_src_url, "-o", "source.tar.gz")
	err = organizetool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	organizetool_cmd_direct := exec.Command("./binary")
	err = organizetool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: openjpeg")
	exec.Command("latte", "install", "openjpeg").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
