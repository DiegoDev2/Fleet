package main

// PythonMatplotlib - Python library for creating static, animated, and interactive visualizations
// Homepage: https://matplotlib.org/

import (
	"fmt"
	
	"os/exec"
)

func installPythonMatplotlib() {
	// Método 1: Descargar y extraer .tar.gz
	pythonmatplotlib_tar_url := "https://files.pythonhosted.org/packages/9e/d8/3d7f706c69e024d4287c1110d74f7dabac91d9843b99eadc90de9efc8869/matplotlib-3.9.2.tar.gz"
	pythonmatplotlib_cmd_tar := exec.Command("curl", "-L", pythonmatplotlib_tar_url, "-o", "package.tar.gz")
	err := pythonmatplotlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonmatplotlib_zip_url := "https://files.pythonhosted.org/packages/9e/d8/3d7f706c69e024d4287c1110d74f7dabac91d9843b99eadc90de9efc8869/matplotlib-3.9.2.zip"
	pythonmatplotlib_cmd_zip := exec.Command("curl", "-L", pythonmatplotlib_zip_url, "-o", "package.zip")
	err = pythonmatplotlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonmatplotlib_bin_url := "https://files.pythonhosted.org/packages/9e/d8/3d7f706c69e024d4287c1110d74f7dabac91d9843b99eadc90de9efc8869/matplotlib-3.9.2.bin"
	pythonmatplotlib_cmd_bin := exec.Command("curl", "-L", pythonmatplotlib_bin_url, "-o", "binary.bin")
	err = pythonmatplotlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonmatplotlib_src_url := "https://files.pythonhosted.org/packages/9e/d8/3d7f706c69e024d4287c1110d74f7dabac91d9843b99eadc90de9efc8869/matplotlib-3.9.2.src.tar.gz"
	pythonmatplotlib_cmd_src := exec.Command("curl", "-L", pythonmatplotlib_src_url, "-o", "source.tar.gz")
	err = pythonmatplotlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonmatplotlib_cmd_direct := exec.Command("./binary")
	err = pythonmatplotlib_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: qhull")
exec.Command("latte", "install", "qhull")
	fmt.Println("Instalando dependencia: patchelf")
exec.Command("latte", "install", "patchelf")
}
