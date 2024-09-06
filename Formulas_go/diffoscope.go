package main

// Diffoscope - In-depth comparison of files, archives, and directories
// Homepage: https://diffoscope.org

import (
	"fmt"
	
	"os/exec"
)

func installDiffoscope() {
	// Método 1: Descargar y extraer .tar.gz
	diffoscope_tar_url := "https://files.pythonhosted.org/packages/9d/54/ccf85796762417128d65cff7656256d9302de155cd9501030f3ff63f7dd5/diffoscope-277.tar.gz"
	diffoscope_cmd_tar := exec.Command("curl", "-L", diffoscope_tar_url, "-o", "package.tar.gz")
	err := diffoscope_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diffoscope_zip_url := "https://files.pythonhosted.org/packages/9d/54/ccf85796762417128d65cff7656256d9302de155cd9501030f3ff63f7dd5/diffoscope-277.zip"
	diffoscope_cmd_zip := exec.Command("curl", "-L", diffoscope_zip_url, "-o", "package.zip")
	err = diffoscope_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diffoscope_bin_url := "https://files.pythonhosted.org/packages/9d/54/ccf85796762417128d65cff7656256d9302de155cd9501030f3ff63f7dd5/diffoscope-277.bin"
	diffoscope_cmd_bin := exec.Command("curl", "-L", diffoscope_bin_url, "-o", "binary.bin")
	err = diffoscope_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diffoscope_src_url := "https://files.pythonhosted.org/packages/9d/54/ccf85796762417128d65cff7656256d9302de155cd9501030f3ff63f7dd5/diffoscope-277.src.tar.gz"
	diffoscope_cmd_src := exec.Command("curl", "-L", diffoscope_src_url, "-o", "source.tar.gz")
	err = diffoscope_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diffoscope_cmd_direct := exec.Command("./binary")
	err = diffoscope_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libmagic")
exec.Command("latte", "install", "libmagic")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
