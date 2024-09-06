package main

// Abcde - Better CD Encoder
// Homepage: https://abcde.einval.com

import (
	"fmt"
	
	"os/exec"
)

func installAbcde() {
	// Método 1: Descargar y extraer .tar.gz
	abcde_tar_url := "https://abcde.einval.com/download/abcde-2.9.3.tar.gz"
	abcde_cmd_tar := exec.Command("curl", "-L", abcde_tar_url, "-o", "package.tar.gz")
	err := abcde_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abcde_zip_url := "https://abcde.einval.com/download/abcde-2.9.3.zip"
	abcde_cmd_zip := exec.Command("curl", "-L", abcde_zip_url, "-o", "package.zip")
	err = abcde_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abcde_bin_url := "https://abcde.einval.com/download/abcde-2.9.3.bin"
	abcde_cmd_bin := exec.Command("curl", "-L", abcde_bin_url, "-o", "binary.bin")
	err = abcde_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abcde_src_url := "https://abcde.einval.com/download/abcde-2.9.3.src.tar.gz"
	abcde_cmd_src := exec.Command("curl", "-L", abcde_src_url, "-o", "source.tar.gz")
	err = abcde_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abcde_cmd_direct := exec.Command("./binary")
	err = abcde_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cdrtools")
exec.Command("latte", "install", "cdrtools")
	fmt.Println("Instalando dependencia: eye-d3")
exec.Command("latte", "install", "eye-d3")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: glyr")
exec.Command("latte", "install", "glyr")
	fmt.Println("Instalando dependencia: lame")
exec.Command("latte", "install", "lame")
	fmt.Println("Instalando dependencia: libdiscid")
exec.Command("latte", "install", "libdiscid")
	fmt.Println("Instalando dependencia: mkcue")
exec.Command("latte", "install", "mkcue")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
	fmt.Println("Instalando dependencia: vorbis-tools")
exec.Command("latte", "install", "vorbis-tools")
}
