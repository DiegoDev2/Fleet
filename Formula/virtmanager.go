package main

// VirtManager - App for managing virtual machines
// Homepage: https://virt-manager.org/

import (
	"fmt"
	
	"os/exec"
)

func installVirtManager() {
	// Método 1: Descargar y extraer .tar.gz
	virtmanager_tar_url := "https://releases.pagure.org/virt-manager/virt-manager-4.1.0.tar.gz"
	virtmanager_cmd_tar := exec.Command("curl", "-L", virtmanager_tar_url, "-o", "package.tar.gz")
	err := virtmanager_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	virtmanager_zip_url := "https://releases.pagure.org/virt-manager/virt-manager-4.1.0.zip"
	virtmanager_cmd_zip := exec.Command("curl", "-L", virtmanager_zip_url, "-o", "package.zip")
	err = virtmanager_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	virtmanager_bin_url := "https://releases.pagure.org/virt-manager/virt-manager-4.1.0.bin"
	virtmanager_cmd_bin := exec.Command("curl", "-L", virtmanager_bin_url, "-o", "binary.bin")
	err = virtmanager_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	virtmanager_src_url := "https://releases.pagure.org/virt-manager/virt-manager-4.1.0.src.tar.gz"
	virtmanager_cmd_src := exec.Command("curl", "-L", virtmanager_src_url, "-o", "source.tar.gz")
	err = virtmanager_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	virtmanager_cmd_direct := exec.Command("./binary")
	err = virtmanager_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docutils")
	exec.Command("latte", "install", "docutils").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
	exec.Command("latte", "install", "adwaita-icon-theme").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cpio")
	exec.Command("latte", "install", "cpio").Run()
	fmt.Println("Instalando dependencia: gtk-vnc")
	exec.Command("latte", "install", "gtk-vnc").Run()
	fmt.Println("Instalando dependencia: gtksourceview4")
	exec.Command("latte", "install", "gtksourceview4").Run()
	fmt.Println("Instalando dependencia: libosinfo")
	exec.Command("latte", "install", "libosinfo").Run()
	fmt.Println("Instalando dependencia: libvirt-glib")
	exec.Command("latte", "install", "libvirt-glib").Run()
	fmt.Println("Instalando dependencia: libvirt-python")
	exec.Command("latte", "install", "libvirt-python").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: osinfo-db")
	exec.Command("latte", "install", "osinfo-db").Run()
	fmt.Println("Instalando dependencia: py3cairo")
	exec.Command("latte", "install", "py3cairo").Run()
	fmt.Println("Instalando dependencia: pygobject3")
	exec.Command("latte", "install", "pygobject3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: spice-gtk")
	exec.Command("latte", "install", "spice-gtk").Run()
	fmt.Println("Instalando dependencia: vte3")
	exec.Command("latte", "install", "vte3").Run()
}
