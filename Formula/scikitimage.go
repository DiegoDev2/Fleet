package main

// ScikitImage - Image processing in Python
// Homepage: https://scikit-image.org

import (
	"fmt"
	
	"os/exec"
)

func installScikitImage() {
	// Método 1: Descargar y extraer .tar.gz
	scikitimage_tar_url := "https://files.pythonhosted.org/packages/5d/c5/bcd66bf5aae5587d3b4b69c74bee30889c46c9778e858942ce93a030e1f3/scikit_image-0.24.0.tar.gz"
	scikitimage_cmd_tar := exec.Command("curl", "-L", scikitimage_tar_url, "-o", "package.tar.gz")
	err := scikitimage_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scikitimage_zip_url := "https://files.pythonhosted.org/packages/5d/c5/bcd66bf5aae5587d3b4b69c74bee30889c46c9778e858942ce93a030e1f3/scikit_image-0.24.0.zip"
	scikitimage_cmd_zip := exec.Command("curl", "-L", scikitimage_zip_url, "-o", "package.zip")
	err = scikitimage_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scikitimage_bin_url := "https://files.pythonhosted.org/packages/5d/c5/bcd66bf5aae5587d3b4b69c74bee30889c46c9778e858942ce93a030e1f3/scikit_image-0.24.0.bin"
	scikitimage_cmd_bin := exec.Command("curl", "-L", scikitimage_bin_url, "-o", "binary.bin")
	err = scikitimage_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scikitimage_src_url := "https://files.pythonhosted.org/packages/5d/c5/bcd66bf5aae5587d3b4b69c74bee30889c46c9778e858942ce93a030e1f3/scikit_image-0.24.0.src.tar.gz"
	scikitimage_cmd_src := exec.Command("curl", "-L", scikitimage_src_url, "-o", "source.tar.gz")
	err = scikitimage_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scikitimage_cmd_direct := exec.Command("./binary")
	err = scikitimage_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: pillow")
	exec.Command("latte", "install", "pillow").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: scipy")
	exec.Command("latte", "install", "scipy").Run()
	fmt.Println("Instalando dependencia: patchelf")
	exec.Command("latte", "install", "patchelf").Run()
}
