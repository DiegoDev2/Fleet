package main

// Crun - Fast and lightweight fully featured OCI runtime and C library
// Homepage: https://github.com/containers/crun

import (
	"fmt"
	
	"os/exec"
)

func installCrun() {
	// Método 1: Descargar y extraer .tar.gz
	crun_tar_url := "https://github.com/containers/crun/releases/download/1.16.1/crun-1.16.1.tar.zst"
	crun_cmd_tar := exec.Command("curl", "-L", crun_tar_url, "-o", "package.tar.gz")
	err := crun_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crun_zip_url := "https://github.com/containers/crun/releases/download/1.16.1/crun-1.16.1.tar.zst"
	crun_cmd_zip := exec.Command("curl", "-L", crun_zip_url, "-o", "package.zip")
	err = crun_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crun_bin_url := "https://github.com/containers/crun/releases/download/1.16.1/crun-1.16.1.tar.zst"
	crun_cmd_bin := exec.Command("curl", "-L", crun_bin_url, "-o", "binary.bin")
	err = crun_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crun_src_url := "https://github.com/containers/crun/releases/download/1.16.1/crun-1.16.1.tar.zst"
	crun_cmd_src := exec.Command("curl", "-L", crun_src_url, "-o", "source.tar.gz")
	err = crun_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crun_cmd_direct := exec.Command("./binary")
	err = crun_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: go-md2man")
exec.Command("latte", "install", "go-md2man")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: libcap")
exec.Command("latte", "install", "libcap")
	fmt.Println("Instalando dependencia: libseccomp")
exec.Command("latte", "install", "libseccomp")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
	fmt.Println("Instalando dependencia: yajl")
exec.Command("latte", "install", "yajl")
}
