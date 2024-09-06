package main

// LibvirtPython - Libvirt virtualization API python binding
// Homepage: https://www.libvirt.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibvirtPython() {
	// Método 1: Descargar y extraer .tar.gz
	libvirtpython_tar_url := "https://download.libvirt.org/python/libvirt-python-10.7.0.tar.gz"
	libvirtpython_cmd_tar := exec.Command("curl", "-L", libvirtpython_tar_url, "-o", "package.tar.gz")
	err := libvirtpython_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvirtpython_zip_url := "https://download.libvirt.org/python/libvirt-python-10.7.0.zip"
	libvirtpython_cmd_zip := exec.Command("curl", "-L", libvirtpython_zip_url, "-o", "package.zip")
	err = libvirtpython_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvirtpython_bin_url := "https://download.libvirt.org/python/libvirt-python-10.7.0.bin"
	libvirtpython_cmd_bin := exec.Command("curl", "-L", libvirtpython_bin_url, "-o", "binary.bin")
	err = libvirtpython_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvirtpython_src_url := "https://download.libvirt.org/python/libvirt-python-10.7.0.src.tar.gz"
	libvirtpython_cmd_src := exec.Command("curl", "-L", libvirtpython_src_url, "-o", "source.tar.gz")
	err = libvirtpython_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvirtpython_cmd_direct := exec.Command("./binary")
	err = libvirtpython_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libvirt")
exec.Command("latte", "install", "libvirt")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
