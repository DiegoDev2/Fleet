package main

// Scipy - Software for mathematics, science, and engineering
// Homepage: https://www.scipy.org

import (
	"fmt"
	
	"os/exec"
)

func installScipy() {
	// Método 1: Descargar y extraer .tar.gz
	scipy_tar_url := "https://files.pythonhosted.org/packages/62/11/4d44a1f274e002784e4dbdb81e0ea96d2de2d1045b2132d5af62cc31fd28/scipy-1.14.1.tar.gz"
	scipy_cmd_tar := exec.Command("curl", "-L", scipy_tar_url, "-o", "package.tar.gz")
	err := scipy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scipy_zip_url := "https://files.pythonhosted.org/packages/62/11/4d44a1f274e002784e4dbdb81e0ea96d2de2d1045b2132d5af62cc31fd28/scipy-1.14.1.zip"
	scipy_cmd_zip := exec.Command("curl", "-L", scipy_zip_url, "-o", "package.zip")
	err = scipy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scipy_bin_url := "https://files.pythonhosted.org/packages/62/11/4d44a1f274e002784e4dbdb81e0ea96d2de2d1045b2132d5af62cc31fd28/scipy-1.14.1.bin"
	scipy_cmd_bin := exec.Command("curl", "-L", scipy_bin_url, "-o", "binary.bin")
	err = scipy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scipy_src_url := "https://files.pythonhosted.org/packages/62/11/4d44a1f274e002784e4dbdb81e0ea96d2de2d1045b2132d5af62cc31fd28/scipy-1.14.1.src.tar.gz"
	scipy_cmd_src := exec.Command("curl", "-L", scipy_src_url, "-o", "source.tar.gz")
	err = scipy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scipy_cmd_direct := exec.Command("./binary")
	err = scipy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: xsimd")
exec.Command("latte", "install", "xsimd")
	fmt.Println("Instalando dependencia: patchelf")
exec.Command("latte", "install", "patchelf")
}
