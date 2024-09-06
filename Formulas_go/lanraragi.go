package main

// Lanraragi - Web application for archival and reading of manga/doujinshi
// Homepage: https://github.com/Difegue/LANraragi

import (
	"fmt"
	
	"os/exec"
)

func installLanraragi() {
	// Método 1: Descargar y extraer .tar.gz
	lanraragi_tar_url := "https://github.com/Difegue/LANraragi/archive/refs/tags/v.0.9.21.tar.gz"
	lanraragi_cmd_tar := exec.Command("curl", "-L", lanraragi_tar_url, "-o", "package.tar.gz")
	err := lanraragi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lanraragi_zip_url := "https://github.com/Difegue/LANraragi/archive/refs/tags/v.0.9.21.zip"
	lanraragi_cmd_zip := exec.Command("curl", "-L", lanraragi_zip_url, "-o", "package.zip")
	err = lanraragi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lanraragi_bin_url := "https://github.com/Difegue/LANraragi/archive/refs/tags/v.0.9.21.bin"
	lanraragi_cmd_bin := exec.Command("curl", "-L", lanraragi_bin_url, "-o", "binary.bin")
	err = lanraragi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lanraragi_src_url := "https://github.com/Difegue/LANraragi/archive/refs/tags/v.0.9.21.src.tar.gz"
	lanraragi_cmd_src := exec.Command("curl", "-L", lanraragi_src_url, "-o", "source.tar.gz")
	err = lanraragi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lanraragi_cmd_direct := exec.Command("./binary")
	err = lanraragi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cpanminus")
exec.Command("latte", "install", "cpanminus")
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
	fmt.Println("Instalando dependencia: giflib")
exec.Command("latte", "install", "giflib")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
	fmt.Println("Instalando dependencia: redis")
exec.Command("latte", "install", "redis")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
}
