package main

// Ntfs3g - Read-write NTFS driver for FUSE
// Homepage: https://www.tuxera.com/community/open-source-ntfs-3g/

import (
	"fmt"
	
	"os/exec"
)

func installNtfs3g() {
	// Método 1: Descargar y extraer .tar.gz
	ntfs3g_tar_url := "https://tuxera.com/opensource/ntfs-3g_ntfsprogs-2022.10.3.tgz"
	ntfs3g_cmd_tar := exec.Command("curl", "-L", ntfs3g_tar_url, "-o", "package.tar.gz")
	err := ntfs3g_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ntfs3g_zip_url := "https://tuxera.com/opensource/ntfs-3g_ntfsprogs-2022.10.3.tgz"
	ntfs3g_cmd_zip := exec.Command("curl", "-L", ntfs3g_zip_url, "-o", "package.zip")
	err = ntfs3g_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ntfs3g_bin_url := "https://tuxera.com/opensource/ntfs-3g_ntfsprogs-2022.10.3.tgz"
	ntfs3g_cmd_bin := exec.Command("curl", "-L", ntfs3g_bin_url, "-o", "binary.bin")
	err = ntfs3g_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ntfs3g_src_url := "https://tuxera.com/opensource/ntfs-3g_ntfsprogs-2022.10.3.tgz"
	ntfs3g_cmd_src := exec.Command("curl", "-L", ntfs3g_src_url, "-o", "source.tar.gz")
	err = ntfs3g_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ntfs3g_cmd_direct := exec.Command("./binary")
	err = ntfs3g_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
}
