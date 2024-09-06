package main

// Curlftpfs - Filesystem for accessing FTP hosts based on FUSE and libcurl
// Homepage: https://curlftpfs.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCurlftpfs() {
	// Método 1: Descargar y extraer .tar.gz
	curlftpfs_tar_url := "https://downloads.sourceforge.net/project/curlftpfs/curlftpfs/0.9.2/curlftpfs-0.9.2.tar.gz"
	curlftpfs_cmd_tar := exec.Command("curl", "-L", curlftpfs_tar_url, "-o", "package.tar.gz")
	err := curlftpfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	curlftpfs_zip_url := "https://downloads.sourceforge.net/project/curlftpfs/curlftpfs/0.9.2/curlftpfs-0.9.2.zip"
	curlftpfs_cmd_zip := exec.Command("curl", "-L", curlftpfs_zip_url, "-o", "package.zip")
	err = curlftpfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	curlftpfs_bin_url := "https://downloads.sourceforge.net/project/curlftpfs/curlftpfs/0.9.2/curlftpfs-0.9.2.bin"
	curlftpfs_cmd_bin := exec.Command("curl", "-L", curlftpfs_bin_url, "-o", "binary.bin")
	err = curlftpfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	curlftpfs_src_url := "https://downloads.sourceforge.net/project/curlftpfs/curlftpfs/0.9.2/curlftpfs-0.9.2.src.tar.gz"
	curlftpfs_cmd_src := exec.Command("curl", "-L", curlftpfs_src_url, "-o", "source.tar.gz")
	err = curlftpfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	curlftpfs_cmd_direct := exec.Command("./binary")
	err = curlftpfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: curl")
exec.Command("latte", "install", "curl")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libfuse@2")
exec.Command("latte", "install", "libfuse@2")
}
