package main

// Fdroidserver - Create and manage Android app repositories for F-Droid
// Homepage: https://f-droid.org

import (
	"fmt"
	
	"os/exec"
)

func installFdroidserver() {
	// Método 1: Descargar y extraer .tar.gz
	fdroidserver_tar_url := "https://files.pythonhosted.org/packages/3c/39/16a78b07797a6fb7fdce85aaa0bb71fb1e459dcdd73ee70be5ff15711059/fdroidserver-2.2.2.tar.gz"
	fdroidserver_cmd_tar := exec.Command("curl", "-L", fdroidserver_tar_url, "-o", "package.tar.gz")
	err := fdroidserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fdroidserver_zip_url := "https://files.pythonhosted.org/packages/3c/39/16a78b07797a6fb7fdce85aaa0bb71fb1e459dcdd73ee70be5ff15711059/fdroidserver-2.2.2.zip"
	fdroidserver_cmd_zip := exec.Command("curl", "-L", fdroidserver_zip_url, "-o", "package.zip")
	err = fdroidserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fdroidserver_bin_url := "https://files.pythonhosted.org/packages/3c/39/16a78b07797a6fb7fdce85aaa0bb71fb1e459dcdd73ee70be5ff15711059/fdroidserver-2.2.2.bin"
	fdroidserver_cmd_bin := exec.Command("curl", "-L", fdroidserver_bin_url, "-o", "binary.bin")
	err = fdroidserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fdroidserver_src_url := "https://files.pythonhosted.org/packages/3c/39/16a78b07797a6fb7fdce85aaa0bb71fb1e459dcdd73ee70be5ff15711059/fdroidserver-2.2.2.src.tar.gz"
	fdroidserver_cmd_src := exec.Command("curl", "-L", fdroidserver_src_url, "-o", "source.tar.gz")
	err = fdroidserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fdroidserver_cmd_direct := exec.Command("./binary")
	err = fdroidserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pybind11")
	exec.Command("latte", "install", "pybind11").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: pillow")
	exec.Command("latte", "install", "pillow").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: qhull")
	exec.Command("latte", "install", "qhull").Run()
	fmt.Println("Instalando dependencia: s3cmd")
	exec.Command("latte", "install", "s3cmd").Run()
	fmt.Println("Instalando dependencia: patchelf")
	exec.Command("latte", "install", "patchelf").Run()
}
