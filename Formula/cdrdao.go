package main

// Cdrdao - Record CDs in Disk-At-Once mode
// Homepage: https://cdrdao.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCdrdao() {
	// Método 1: Descargar y extraer .tar.gz
	cdrdao_tar_url := "https://github.com/cdrdao/cdrdao/archive/refs/tags/rel_1_2_5.tar.gz"
	cdrdao_cmd_tar := exec.Command("curl", "-L", cdrdao_tar_url, "-o", "package.tar.gz")
	err := cdrdao_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdrdao_zip_url := "https://github.com/cdrdao/cdrdao/archive/refs/tags/rel_1_2_5.zip"
	cdrdao_cmd_zip := exec.Command("curl", "-L", cdrdao_zip_url, "-o", "package.zip")
	err = cdrdao_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdrdao_bin_url := "https://github.com/cdrdao/cdrdao/archive/refs/tags/rel_1_2_5.bin"
	cdrdao_cmd_bin := exec.Command("curl", "-L", cdrdao_bin_url, "-o", "binary.bin")
	err = cdrdao_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdrdao_src_url := "https://github.com/cdrdao/cdrdao/archive/refs/tags/rel_1_2_5.src.tar.gz"
	cdrdao_cmd_src := exec.Command("curl", "-L", cdrdao_src_url, "-o", "source.tar.gz")
	err = cdrdao_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdrdao_cmd_direct := exec.Command("./binary")
	err = cdrdao_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: lame")
	exec.Command("latte", "install", "lame").Run()
	fmt.Println("Instalando dependencia: libao")
	exec.Command("latte", "install", "libao").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: mad")
	exec.Command("latte", "install", "mad").Run()
}
